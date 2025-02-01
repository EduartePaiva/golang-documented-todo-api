import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { Toaster } from "react-hot-toast";
import App from "./App.tsx";
import BackgroundLayout from "./components/background-layout.tsx";
import TodoProvider from "./context/todo-context/todo-context.tsx";
import UserProvider from "./context/user-context.tsx";
import "./index.css";

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <UserProvider>
            <TodoProvider>
                <BackgroundLayout>
                    <App />
                </BackgroundLayout>
            </TodoProvider>
            <Toaster />
        </UserProvider>
    </StrictMode>
);
