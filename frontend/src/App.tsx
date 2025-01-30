import { Fragment } from "react";
import Header from "./components/header";
import TodoItem from "./components/todo-item";
import { Button } from "./components/ui/button";
import { useTodoContext } from "./context/todo-context";

function App() {
    const { todos, createTodo } = useTodoContext();

    return (
        <div className="">
            <Header />
            <div className="relative h-36 w-full">
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
