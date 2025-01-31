import { TodoItemType } from "@/types/todo-type";
import React, { createContext, useContext, useRef, useState } from "react";
import {
    createTodo,
    deleteTodo,
    getTodosFromLocalStorage,
    scheduleTextUpdate,
    updateTodo,
} from "./todo-context-service";

type TodoProviderProps = { children: React.ReactNode };
type TodoContextType = {
    todos: TodoItemType[];
    createTodo: () => void;
    deleteTodo: (id: string) => void;
    updateTodo: (params: { id: string; text?: string; done?: boolean }) => void;
    scheduleTextUpdate: (text: string, id: string) => void;
};

const TodoContext = createContext<TodoContextType | null>(null);

export default function TodoProvider({ children }: TodoProviderProps) {
    const lastTimeoutId = useRef(-1);
    const [todos, setTodo] = useState<TodoItemType[]>(getTodosFromLocalStorage);

    return (
        <TodoContext.Provider
            value={{
                todos,
                createTodo: createTodo(setTodo),
                deleteTodo: deleteTodo(setTodo),
                updateTodo: updateTodo(setTodo, lastTimeoutId),
                scheduleTextUpdate: scheduleTextUpdate(setTodo, lastTimeoutId),
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
