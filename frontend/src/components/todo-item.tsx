import { useTodoContext } from "@/context/todo-context/todo-context";
import { useUserContext } from "@/context/user-context";
import { TodoItemType } from "@/types/todo-type";
import { X } from "lucide-react";
import toast from "react-hot-toast";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { Textarea } from "./ui/textarea";

export default function TodoItem({ text, id, done }: TodoItemType) {
    const { deleteTodo, updateTodo, scheduleTextUpdate } = useTodoContext();
    const { user } = useUserContext();
    return (
        <div className="flex w-min flex-col justify-center gap-1 rounded-sm bg-slate-50/50 p-2">
            <div className="flex items-center justify-between">
                <div>
                    <label htmlFor={id} className="flex items-center gap-2">
                        <Checkbox
                            id={id}
                            className="h-5 w-5"
                            defaultChecked={done}
                            onCheckedChange={(value) => {
                                updateTodo({
                                    id,
                                    done: value !== "indeterminate" && value,
                                });
                            }}
                        />
                        <span className="select-none font-semibold hover:cursor-pointer">
                            Done
                        </span>
                    </label>
                </div>
                <Button
                    size={"icon"}
                    variant={"destructive"}
                    className="h-7 w-7"
                    onClick={() => {
                        toast.dismiss();
                        const toastId = toast.loading("Deleting...");
                        deleteTodo(id, user.loggedIn)
                            .then((statusCode) => {
                                switch (statusCode) {
                                    case -1:
                                        toast.success("Todo deleted locally", {
                                            id: toastId,
                                            icon: "⚠",
                                        });
                                        break;
                                    case 204:
                                        toast.success(
                                            "Todo deleted locally and in the database",
                                            { id: toastId }
                                        );
                                        break;
                                    default:
                                        toast.error(
                                            `Some unexpected error happened, status code: ${statusCode}`,
                                            { id: toastId }
                                        );
                                }
                            })
                            .catch(() =>
                                toast.error("Network error", { id: toastId })
                            );
                    }}
                >
                    <X />
                </Button>
            </div>
            <Textarea
                defaultValue={text}
                onChange={(e) => scheduleTextUpdate(e.target.value, id)}
                onBlur={(e) => updateTodo({ id, text: e.target.value })}
                className="w-[17rem] sm:w-[25rem] md:w-[30rem] lg:w-[35rem]"
            />
        </div>
    );
}
