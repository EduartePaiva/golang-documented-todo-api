import { Fragment } from "react";
import LoginBtn from "./components/login-btn";
import TodoItem from "./components/todo-item";
import { Button } from "./components/ui/button";
import { useTodoContext } from "./context/todo-context";

function App() {
    const { todos, createTodo } = useTodoContext();

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
                    onClick={createTodo}
                >
                    Create Todo
                </Button>
            </div>
            <div className="flex flex-col items-center gap-2">
                {todos.map((todo) => (
                    <Fragment key={todo.id}>
                        <TodoItem {...todo} />
                    </Fragment>
                ))}
            </div>
        </div>
    );
}

export default App;
