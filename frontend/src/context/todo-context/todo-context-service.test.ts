// test update todo
// when it's called the updatedAt field must be updated to a recent date

import { TodoItemType } from "@/types/todo-type";
import { describe, test, vi } from "vitest";
import { updateTodo } from "./todo-context-service";

describe("test update todo", () => {
    test("when it's called the updatedAt field must be updated to a recent date", async () => {
        const todo: TodoItemType = {
            id: crypto.randomUUID(),
            text: "type something...",
            done: false,
            createdAt: new Date().toJSON(),
            updatedAt: new Date().toJSON(),
        };
        let someState = [todo];

        // basically what set todo does:
        // it receives a callback and then call this callback with the prev value
        // then it receives the new value and assign this value to todos
        const setTodoMock = vi.fn().mockImplementation((action) => {
            someState =
                typeof action === "function" ? action(someState) : action;
        });

        const lastTimeoutMock: React.MutableRefObject<number> = { current: -1 };

        const update = updateTodo(setTodoMock, lastTimeoutMock);
    });
});
