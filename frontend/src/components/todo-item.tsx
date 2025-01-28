import { TodoItemType, useTodoContext } from "@/context/todo-context";
import { X } from "lucide-react";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { Textarea } from "./ui/textarea";

export default function TodoItem({ text, createdAt }: TodoItemType) {
    const { deleteTodo } = useTodoContext();
    return (
        <div className="flex w-min flex-col justify-center gap-1 rounded-sm bg-slate-50/50 p-2">
            <div className="flex items-center justify-between">
                <div>
                    <label
                        htmlFor={createdAt}
                        className="flex items-center gap-2"
                    >
                        <Checkbox id={createdAt} className="h-5 w-5" />
                        <span className="select-none font-semibold hover:cursor-pointer">
                            Done
                        </span>
                    </label>
                </div>
                <Button
                    size={"icon"}
                    variant={"destructive"}
                    className="h-7 w-7"
                    onClick={() => deleteTodo(createdAt)}
                >
                    <X />
                </Button>
            </div>
            <Textarea
                defaultValue={text}
                className="w-[17rem] sm:w-[25rem] md:w-[30rem] lg:w-[35rem]"
            />
        </div>
    );
}
