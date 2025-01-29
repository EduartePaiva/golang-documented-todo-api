import React, { createContext, useContext, useState } from "react";
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
    deleteTodo: (createdAt: string) => void;
    updateTodo: (id: string, text?: string, done?: boolean) => void;
};

const TodoContext = createContext<TodoContextType | null>(null);

function setStorage(todos: TodoItemType[]) {
    window.localStorage.setItem("todos", JSON.stringify(todos));
}

export default function TodoProvider({ children }: TodoProviderProps) {
    const [todos, setTodo] = useState<TodoItemType[]>(() => {
        const localTodo = window.localStorage.getItem("todos") as null | string;
        if (localTodo) {
            console.log(JSON.parse(localTodo));
            return JSON.parse(localTodo);
        }
        return [];
    });
    const createTodo = () => {
        setTodo((prev) => {
            prev.push({
                id: crypto.randomUUID(),
                text: "type something...",
                done: false,
                createdAt: new Date().toJSON(),
                updatedAt: new Date().toJSON(),
            });
            setStorage(prev);
            return [...prev];
        });
    };
    const deleteTodo = (id: string) => {
        setTodo((prev) => {
            const idx = prev.findIndex((todo) => todo.id === id);
            prev.splice(idx, 1);
            setStorage(prev);
            return [...prev];
        });
    };

    const updateTodo = (id: string, text?: string, done?: boolean) => {
        if (text === undefined && done === undefined) {
            return;
        }
        setTodo((prev) => {
            const idx = prev.findIndex((todo) => todo.id === id);
            if (text !== undefined) {
                prev[idx].text = text;
            }
            if (done !== undefined) {
                prev[idx].done = done;
            }
            setStorage(prev);
            return [...prev];
        });
    };

    return (
        <TodoContext.Provider
            value={{
                todos,
                createTodo,
                deleteTodo,
                updateTodo,
            }}
        >
            {children}
        </TodoContext.Provider>
    );
}

export function useTodoContext() {
    const context = useContext(TodoContext);
    if (context === null) {
        throw new Error("useTodoContext must be used within a ThemeProvider");
    }
    return context;
}
