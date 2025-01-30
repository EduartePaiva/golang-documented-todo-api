import { useUserContext } from "@/context/user-context";
import { BiTask } from "react-icons/bi";
import LoginBtn from "./login-btn";
import UserAvatar from "./user-avatar-with-options";

export default function Header() {
    const { user } = useUserContext();
    return (
        <header className="flex h-12 w-full items-center justify-between bg-slate-50/40 px-3 shadow-sm">
            <div className="flex items-center gap-2">
                <BiTask size={25} />
                <span className="font-mono text-2xl font-semibold">
                    Todo App
                </span>
            </div>
            {user.loggedIn ? (
                <UserAvatar avatar_url={user.avatar_url} name={user.name} />
            ) : (
                <LoginBtn variant={"outline"} size={"sm"} />
            )}
        </header>
    );
}
