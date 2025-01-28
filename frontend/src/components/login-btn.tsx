import React from "react";
import { Button, ButtonProps } from "./ui/button";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import ProviderLoginBtn from "./provider-button";
import { FaGithub, FaGoogle } from "react-icons/fa";

type LoginBtnProps = ButtonProps & React.RefAttributes<HTMLButtonElement>;

export default function LoginBtn(props: LoginBtnProps) {
    return (
        <DropdownMenu dir="ltr">
            <DropdownMenuTrigger asChild>
                <Button {...props}>Log In</Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
                <DropdownMenuLabel>LogIn with...</DropdownMenuLabel>
                <DropdownMenuSeparator />
                <DropdownMenuItem asChild>
                    <ProviderLoginBtn
                        LogoIcon={FaGoogle}
                        name="Google"
                        href="/api/v1/login/google"
                        className="text-[#db4639] hover:text-[#db4639]/90"
                    />
                </DropdownMenuItem>
                <DropdownMenuItem asChild>
                    <ProviderLoginBtn
                        LogoIcon={FaGithub}
                        name="Github"
                        href="/api/v1/login/github"
                    />
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
