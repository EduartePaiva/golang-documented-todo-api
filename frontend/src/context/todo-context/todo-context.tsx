import {
    findWitchTodosToUpdateRemote,
    syncLocalWithRemoteData,
} from "@/lib/sync-data";
import { TodoItemType } from "@/types/todo-type";
import React, {
    createContext,
    useContext,
    useEffect,
    useRef,
    useState,
} from "react";
import toast from "react-hot-toast";
import { useUserContext } from "../user-context";
import {
    createTodo,
    deleteTodo,
    getTodosFromLocalStorage,
    scheduleTextUpdate,
    setStorage,
    updateTodo,
} from "./todo-context-service";

type TodoProviderProps = { children: React.ReactNode };
type TodoContextType = {
    todos: TodoItemType[];
    createTodo: () => void;
    deleteTodo: (id: string, isLoggedIn: boolean) => Promise<number>;
    updateTodo: (params: { id: string; text?: string; done?: boolean }) => void;
    scheduleTextUpdate: (text: string, id: string) => void;
};

const TodoContext = createContext<TodoContextType | null>(null);

async function syncTodos(): Promise<TodoItemType[]> {
    const result = await fetch("/api/v1/tasks");
    if (!result.ok || result.status != 200) {
        throw new Error("failed to fetch the tasks from database");
    }
    const remote = await result.json();
    const local = getTodosFromLocalStorage();
    const todosToUpdate = findWitchTodosToUpdateRemote(local, remote);
    const updateRes = await fetch("/api/v1/tasks", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(todosToUpdate),
    });
    if (!updateRes.ok || updateRes.status != 200) {
        throw new Error("failed to update the database");
    }
    const finalTodos = syncLocalWithRemoteData(local, remote);
    setStorage(finalTodos);
    return finalTodos;
}

export default function TodoProvider({ children }: TodoProviderProps) {
    const lastTimeoutId = useRef(-1);
    const [todos, setTodo] = useState<TodoItemType[]>(getTodosFromLocalStorage);
    const { user } = useUserContext();

    useEffect(() => {
        if (!user.loggedIn) {
            return;
        }
        const toastId = toast.loading("Syncing with the database");
        syncTodos()
            .then((res) => {
                setTodo(res);
                toast.success("Successfully synced with the database", {
                    id: toastId,
                });
            })
            .catch((err) => {
                console.error(err);
                if (err instanceof Error) {
                    toast.error(err.message, { id: toastId });
                } else {
                    toast.error("Failed to sync with the database", {
                        id: toastId,
                    });
                }
            });
    }, [user]);

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
