import { TodoItemType } from "@/types/todo-type";

/**
 * This function deeply clone the todos
 */
export function syncLocalWithRemoteData(
    local: TodoItemType[],
    remote: TodoItemType[]
): TodoItemType[] {
    if (local.length === 0 && remote.length === 0) {
        return [];
    }
    if (local.length == 0) {
        return remote.map((t) => ({ ...t }));
    }
    if (remote.length == 0) {
        return local.map((t) => ({ ...t }));
    }

    const syncedTodos = new Map<string, TodoItemType>();

    for (const todo of remote) {
        syncedTodos.set(todo.id, { ...todo });
    }
    for (const todo of local) {
        const remoteTodo = syncedTodos.get(todo.id);
        // check if it's already in from remote
        if (remoteTodo !== undefined) {
            const remoteUpdatedAt = new Date(remoteTodo.updatedAt);
            const localUpdatedAt = new Date(todo.updatedAt);
            // if remote is the most recent, we can just skip adding the local
            if (remoteUpdatedAt > localUpdatedAt) {
                continue;
            }
        }
        syncedTodos.set(todo.id, { ...todo });
    }

    return [...syncedTodos.values()];
}

/**
 * this function compares local with remote and return
 * a list of todos in local that are more recent than remote
 */
export function findWitchTodosToUpdateRemote(
    local: TodoItemType[],
    remote: TodoItemType[]
): TodoItemType[] {
    if (local.length === 0) {
        return [];
    }
    if (remote.length === 0) {
        return local;
    }

    const remoteMap = new Map<string, TodoItemType>();
    for (const todo of remote) {
        remoteMap.set(todo.id, todo);
    }

    const todosToUpdate: TodoItemType[] = [];
    for (const todo of local) {
        const rmt = remoteMap.get(todo.id);
        //whatever local have that remote don't
        if (rmt === undefined) {
            todosToUpdate.push(todo);
            continue;
        }

        const remoteUpdatedAt = new Date(rmt.updatedAt);
        const localUpdatedAt = new Date(todo.updatedAt);
        // if local is the most recent, add local
        if (localUpdatedAt > remoteUpdatedAt) {
            todosToUpdate.push(todo);
        }
    }
    return todosToUpdate;
}
