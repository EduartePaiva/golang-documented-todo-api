import { TodoItemType } from "./../../types/todo-type";

type setTodoType = React.Dispatch<React.SetStateAction<TodoItemType[]>>;

export function setStorage(todos: TodoItemType[]) {
    window.localStorage.setItem("todos", JSON.stringify(todos));
}

export function createTodo(setTodo: setTodoType) {
    return () => {
        const newTodo: TodoItemType = {
            id: crypto.randomUUID(),
            text: "type something...",
            done: false,
            createdAt: new Date().toJSON(),
            updatedAt: new Date().toJSON(),
        };

        setTodo((prev) => {
            const current = [newTodo, ...prev];
            setStorage(current);
            return current;
        });
        try {
            fetch("/api/v1/tasks", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify([newTodo]),
            });
        } catch (err) {
            console.error(err);
        }
    };
}

export function deleteTodo(setTodo: setTodoType) {
    /**
     * returns -1 if the user is not logged in, meaning only the local storage was updated
     */
    return async (id: string, isLoggedIn: boolean): Promise<number> => {
        let status = -1;
        if (isLoggedIn) {
            const res = await fetch(`/api/v1/tasks/${id}`, {
                method: "DELETE",
            });
            status = res.status;
        }
        setTodo((prev) => {
            const current = prev.filter((todo) => todo.id !== id);
            setStorage(current);
            return current;
        });
        return status;
    };
}

export function updateTodo(
    setTodo: setTodoType,
    lastTimeoutId: React.MutableRefObject<number>
) {
    return ({
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
        setTodo((prev) => {
            const current = prev.map((todo) => {
                if (todo.id === id) {
                    const newTodo: TodoItemType = {
                        ...todo,
                        updatedAt: new Date().toJSON(),
                    };
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

        try {
            fetch("/api/v1/tasks", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify([
                    {
                        id,
                        done: done,
                    },
                ]),
            });
        } catch (err) {
            console.error(err);
        }
    };
}

export function scheduleTextUpdate(
    setTodo: setTodoType,
    lastTimeoutId: React.MutableRefObject<number>
) {
    const update = updateTodo(setTodo, lastTimeoutId);
    return (text: string, id: string) => {
        // clear the last timeout
        window.clearTimeout(lastTimeoutId.current);
        // save the current timeout id
        lastTimeoutId.current = window.setTimeout(
            () => update({ id, text }),
            3000
        );
    };
}

export function getTodosFromLocalStorage(): TodoItemType[] {
    const localTodo = window.localStorage.getItem("todos") as null | string;
    if (localTodo) {
        return JSON.parse(localTodo);
    }
    return [];
}
