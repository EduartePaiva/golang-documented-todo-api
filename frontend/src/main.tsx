import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import App from "./App.tsx";
import BackgroundLayout from "./components/background-layout.tsx";
import "./index.css";

createRoot(document.getElementById("root")!).render(
    <BrowserRouter>
        <Routes>
            <Route element={<BackgroundLayout />}>
                <Route index element={<App />} />
            </Route>
        </Routes>
    </BrowserRouter>
);
