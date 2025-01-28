import { Fragment, useState } from "react";
import LoginBtn from "./components/login-btn";
import TodoItem from "./components/todo-item";
import { Button } from "./components/ui/button";

export interface TodoItemType {
    text: string;
    done: boolean;
    updatedAt: Date;
    createdAt: Date;
}

function App() {
    const [todos, addTodo] = useState<TodoItemType[]>([]);

    return (
        <div className="">
            <div className="relative h-36 w-full">
                <LoginBtn
                    className="absolute left-[calc(100%-0.75rem)] top-3 -translate-x-[100%]"
                    variant={"outline"}
                    size={"sm"}
                />
                <Button
                    className="absolute left-[50%] top-12 -translate-x-[50%]"
                    onClick={() => {
                        addTodo((prev) => {
                            prev.push({
                                createdAt: new Date(),
                                text: "type something...",
                                done: false,
                                updatedAt: new Date(),
                            });
                            return [...prev];
                        });
                    }}
                >
                    Create Todo
                </Button>
            </div>
            <div className="flex flex-col items-center gap-2">
                {todos.map((todo) => (
                    <Fragment key={todo.createdAt.toISOString()}>
                        <TodoItem {...todo} />
                    </Fragment>
                ))}
            </div>
        </div>
    );
}

export default App;
