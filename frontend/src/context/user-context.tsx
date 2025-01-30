import React, { createContext, useContext, useEffect, useState } from "react";
export interface UserData {
    name: string;
    avatar_url: string;
    loggedIn: boolean;
}

type UserProviderProps = { children: React.ReactNode };
type TodoContextType = {
    user: UserData;
};

const TodoContext = createContext<TodoContextType | null>(null);

async function fetchUser(): Promise<UserData> {
    try {
        const response = await fetch("/api/v1/login/getuser");
        if (response.status != 200) {
            return {
                avatar_url: "",
                name: "",
                loggedIn: false,
            };
        }
        const data = await response.json();
        return {
            avatar_url: data.avatar_url,
            name: data.name,
            loggedIn: true,
        };
    } catch (e) {
        console.error(e);
        return {
            avatar_url: "",
            name: "",
            loggedIn: false,
        };
    }
}

export default function UserProvider({ children }: UserProviderProps) {
    const [user, setUser] = useState<UserData>({
        avatar_url: "",
        name: "",
        loggedIn: false,
    });

    useEffect(() => {
        fetchUser().then((user) => setUser(user));
    }, []);

    return (
        <TodoContext.Provider
            value={{
                user,
            }}
        >
            {children}
        </TodoContext.Provider>
    );
}

// eslint-disable-next-line react-refresh/only-export-components
export function useUserContext() {
    const context = useContext(TodoContext);
    if (context === null) {
        throw new Error("useTodoContext must be used within a ThemeProvider");
    }
    return context;
}
