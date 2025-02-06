import { TodoItemType } from "@/types/todo-type";
import { randomUUID } from "crypto";
import { describe, expect, it } from "vitest";
import {
    findWitchTodosToUpdateRemote,
    syncLocalWithRemoteData,
} from "./sync-data";

describe("Test the syncLocalWithRemote function", () => {
    it("Two empty input return a empty array", () => {
        expect(syncLocalWithRemoteData([], [])).toStrictEqual([]);
    });

    it("If some local contains values that remote don't, it'll return local", () => {
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        const result: TodoItemType[] = [{ ...local[0] }];
        expect(syncLocalWithRemoteData(local, [])).toStrictEqual(result);
    });
    it("If some remote contains values that local don't, it'll return remote", () => {
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        const result: TodoItemType[] = [{ ...local[0] }];
        expect(syncLocalWithRemoteData(local, [])).toStrictEqual(result);
    });

    it("If both are exactly the same, it'll just return one of them", () => {
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        const remote: TodoItemType[] = [{ ...local[0] }];
        const result: TodoItemType[] = [{ ...local[0] }];
        expect(syncLocalWithRemoteData(local, remote)).toStrictEqual(result);
    });
    it("if the updatedAt of local is more recent, it'll return local", () => {
        const local: TodoItemType[] = [];
        const updatedAt = new Date();
        updatedAt.setDate(updatedAt.getDate() + 3);
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: updatedAt.toJSON(),
        });
        const remote: TodoItemType[] = [
            { ...local[0], updatedAt: new Date().toJSON() },
        ];
        const result: TodoItemType[] = [{ ...local[0] }];
        expect(syncLocalWithRemoteData(local, remote)).toStrictEqual(result);
    });
    it("if the updatedAt of remote is more recent, it'll return remote", () => {
        const local: TodoItemType[] = [];
        const updatedAt = new Date();
        updatedAt.setDate(updatedAt.getDate() + 3);
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        const remote: TodoItemType[] = [
            { ...local[0], updatedAt: updatedAt.toJSON() },
        ];
        const result: TodoItemType[] = [{ ...remote[0] }];
        expect(syncLocalWithRemoteData(local, remote)).toStrictEqual(result);
    });
});

describe("Test findWitchTodosToUpdateRemote", () => {
    it("if both are empty", () => {
        expect(findWitchTodosToUpdateRemote([], [])).toStrictEqual([]);
    });
    it("if local have some value but remote is empty", () => {
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        expect(findWitchTodosToUpdateRemote(local, [])).toStrictEqual(local);
    });

    it("if remote have some value but local is empty", () => {
        const remote: TodoItemType[] = [];
        remote.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        expect(findWitchTodosToUpdateRemote([], remote)).toStrictEqual([]);
    });
    it("if both have different values", () => {
        const remote: TodoItemType[] = [];
        remote.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        expect(findWitchTodosToUpdateRemote(local, remote)).toStrictEqual(
            local
        );
    });
    it("if both have same value but remote is more recent", () => {
        const updatedAt = new Date();
        updatedAt.setDate(updatedAt.getDate() + 3);
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: new Date().toJSON(),
        });
        const remote = [{ ...local[0], updatedAt: updatedAt.toJSON() }];
        expect(findWitchTodosToUpdateRemote(local, remote)).toStrictEqual([]);
    });
    it("if both have same value but local is more recent", () => {
        const updatedAt = new Date();
        updatedAt.setDate(updatedAt.getDate() + 3);
        const local: TodoItemType[] = [];
        local.push({
            id: randomUUID(),
            createdAt: new Date().toJSON(),
            done: false,
            text: "some value",
            updatedAt: updatedAt.toJSON(),
        });
        const remote = [{ ...local[0], updatedAt: new Date().toJSON() }];
        expect(findWitchTodosToUpdateRemote(local, remote)).toStrictEqual(
            local
        );
    });

    it("should be empty because both are equal", () => {
        const local = [
            {
                id: "78165d95-5a65-4abf-807b-55d88c44fa39",
                text: "first",
                done: false,
                createdAt: "2025-02-06T10:10:51.427Z",
                updatedAt: "2025-02-06T10:13:52.222Z",
            },
            {
                id: "fb284925-11b6-49d2-b8d4-220e7eca8b2d",
                text: "second",
                done: false,
                createdAt: "2025-02-06T10:10:50.117Z",
                updatedAt: "2025-02-06T10:13:53.996Z",
            },
            {
                id: "3e6714b3-a727-42ad-a5ab-d858ee7fddb9",
                text: "third",
                done: false,
                createdAt: "2025-02-06T10:10:49.403Z",
                updatedAt: "2025-02-06T10:13:56.058Z",
            },
            {
                id: "fffb9c3b-2db8-46b0-917d-afc2c546751d",
                text: "fourth",
                done: false,
                updatedAt: "2025-02-06T10:13:57.213Z",
                createdAt: "2025-02-05T17:30:20.579Z",
            },
        ];
        const remote = [
            {
                id: "78165d95-5a65-4abf-807b-55d88c44fa39",
                text: "first",
                done: false,
                updatedAt: "2025-02-06T10:10:51.430738Z",
                createdAt: "2025-02-06T10:13:52.222Z",
            },
            {
                id: "fb284925-11b6-49d2-b8d4-220e7eca8b2d",
                text: "second",
                done: false,
                updatedAt: "2025-02-06T10:10:50.120259Z",
                createdAt: "2025-02-06T10:13:53.996Z",
            },
            {
                id: "3e6714b3-a727-42ad-a5ab-d858ee7fddb9",
                text: "third",
                done: false,
                updatedAt: "2025-02-06T10:10:49.407286Z",
                createdAt: "2025-02-06T10:13:56.058Z",
            },
            {
                id: "fffb9c3b-2db8-46b0-917d-afc2c546751d",
                text: "fourth",
                done: false,
                updatedAt: "2025-02-05T17:30:20.582204Z",
                createdAt: "2025-02-06T10:13:57.213Z",
            },
        ];

        expect(findWitchTodosToUpdateRemote(local, remote)).toStrictEqual([]);
    });
});
