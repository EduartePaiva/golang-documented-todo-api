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
});
