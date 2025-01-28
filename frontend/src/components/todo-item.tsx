import { TodoItemType } from "@/App";
import { Textarea } from "./ui/textarea";
import { X } from "lucide-react";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";

export default function TodoItem({ text, createdAt }: TodoItemType) {
    return (
        <div className="flex w-min flex-col justify-center gap-1 rounded-sm bg-slate-50/50 p-2">
            <div className="flex items-center justify-between">
                <div>
                    <label
                        htmlFor={createdAt.toISOString()}
                        className="flex items-center gap-2"
                    >
                        <Checkbox
                            id={createdAt.toISOString()}
                            className="h-5 w-5"
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
                >
                    <X />
                </Button>
            </div>
            <Textarea
                value={text}
                className="w-[17rem] sm:w-[25rem] md:w-[30rem] lg:w-[35rem]"
            />
        </div>
    );
}
