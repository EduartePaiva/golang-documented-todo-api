import React, { createContext, useContext, useRef, useState } from "react";
export interface TodoItemType {
    id: string;
    text: string;
    done: boolean;
    updatedAt: string;
    createdAt: string;
}

type TodoProviderProps = { children: React.ReactNode };
type TodoContextType = {
    todos: TodoItemType[];
    createTodo: () => void;
    deleteTodo: (id: string) => void;
    updateTodo: (params: { id: string; text?: string; done?: boolean }) => void;
    scheduleTextUpdate: (text: string, id: string) => void;
};

const TodoContext = createContext<TodoContextType | null>(null);

function setStorage(todos: TodoItemType[]) {
    window.localStorage.setItem("todos", JSON.stringify(todos));
}

export default function TodoProvider({ children }: TodoProviderProps) {
    const lastTimeoutId = useRef(-1);

    const [todos, setTodo] = useState<TodoItemType[]>(() => {
        const localTodo = window.localStorage.getItem("todos") as null | string;
        if (localTodo) {
            return JSON.parse(localTodo);
        }
        return [];
    });
    const createTodo = () => {
        setTodo((prev) => {
            const current = [
                ...prev,
                {
                    id: crypto.randomUUID(),
                    text: "type something...",
                    done: false,
                    createdAt: new Date().toJSON(),
                    updatedAt: new Date().toJSON(),
                },
            ];

            setStorage(current);
            return current;
        });
    };
    const deleteTodo = (id: string) => {
        setTodo((prev) => {
            const current = prev.filter((todo) => todo.id !== id);
            setStorage(current);
            return current;
        });
    };

    const updateTodo = ({
        id,
        done,
        text,
    }: {
        id: string;
        text?: string;
        done?: boolean;
    }) => {
        if (text === undefined && done === undefined) {
            return;
        }
        // this line below clear last timeout if it wasn't cleaned already
        window.clearTimeout(lastTimeoutId.current);
        console.log("saving todo at timeout: ", lastTimeoutId.current);
        setTodo((prev) => {
            const current = prev.map((todo) => {
                if (todo.id === id) {
                    const newTodo: TodoItemType = { ...todo };
                    if (text !== undefined) {
                        newTodo.text = text;
                    }
                    if (done !== undefined) {
                        newTodo.done = done;
                    }
                    return newTodo;
                }
                return todo;
            });

            setStorage(current);
            return current;
        });
    };

    const scheduleTextUpdate = (text: string, id: string) => {
        // clear the last timeout
        window.clearTimeout(lastTimeoutId.current);
        // save the current timeout id
        lastTimeoutId.current = window.setTimeout(
            () => updateTodo({ id, text }),
            3000
        );
    };

    return (
        <TodoContext.Provider
            value={{
                todos,
                createTodo,
                deleteTodo,
                updateTodo,
                scheduleTextUpdate,
            }}
        >
            {children}
        </TodoContext.Provider>
    );
}

// eslint-disable-next-line react-refresh/only-export-components
export function useTodoContext() {
    const context = useContext(TodoContext);
    if (context === null) {
        throw new Error("useTodoContext must be used within a ThemeProvider");
    }
    return context;
}
