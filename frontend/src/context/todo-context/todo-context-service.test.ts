import { TodoItemType } from "@/types/todo-type";
import { afterEach, beforeEach, describe, expect, test, vi } from "vitest";
import { updateTodo } from "./todo-context-service";

describe("test update todo", () => {
    beforeEach(() => vi.useFakeTimers());
    afterEach(() => {
        vi.useRealTimers();
        vi.clearAllMocks();
    });

    test("when it's called the updatedAt field must be updated to a recent date, and the text be the new text", async () => {
        const todo: TodoItemType = {
            id: crypto.randomUUID(),
            text: "type something...",
            done: false,
            createdAt: new Date().toJSON(),
            updatedAt: new Date("2021").toJSON(),
        };
        let someState = [todo];
        // basically what set todo does:
        // it receives a callback and then call this callback with the prev value
        // then it receives the new value and assign this value to todos
        const setTodoMock = vi.fn().mockImplementation((action) => {
            someState =
                typeof action === "function" ? action(someState) : action;
        });
        const putDatabase = async () => {};

        const lastTimeoutMock: React.MutableRefObject<number> = { current: -1 };
        const update = updateTodo(setTodoMock, putDatabase, lastTimeoutMock);

        const fakeDate = new Date();
        vi.setSystemTime(fakeDate);
        update({ id: todo.id, text: "new text" });
        expect(someState[0]).toMatchObject({
            text: "new text",
            updatedAt: fakeDate.toJSON(),
        });
    });
});
