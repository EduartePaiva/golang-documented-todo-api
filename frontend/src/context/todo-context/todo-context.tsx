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
    deleteTodo: (id: string) => void;
    updateTodo: (params: { id: string; text?: string; done?: boolean }) => void;
    scheduleTextUpdate: (text: string, id: string) => void;
};

const TodoContext = createContext<TodoContextType | null>(null);

export default function TodoProvider({ children }: TodoProviderProps) {
    const lastTimeoutId = useRef(-1);
    const [todos, setTodo] = useState<TodoItemType[]>(getTodosFromLocalStorage);
    const { user } = useUserContext();

    useEffect(() => {
        if (!user.loggedIn) {
            return;
        }
        // get the todos from the database
        const toastId = toast.loading("Syncing with the database");
        fetch("/api/v1/tasks")
            .then((res) => {
                if (res.ok && res.status == 200) {
                    return res.json();
                } else {
                    throw new Error("failed to fetch database");
                }
            })
            .then((remote) => {
                const local = getTodosFromLocalStorage();
                const todosToUpdate = findWitchTodosToUpdateRemote(
                    local,
                    remote
                );
                console.log(todosToUpdate);
                const finalTodos = syncLocalWithRemoteData(local, remote);
                setStorage(finalTodos);
                setTodo(finalTodos);
                toast.success("Your todos was synced successfully", {
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
