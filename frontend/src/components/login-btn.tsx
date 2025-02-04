import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import React from "react";
import { FaGithub, FaGoogle } from "react-icons/fa";
import ProviderLoginBtn from "./provider-button";
import { Button, ButtonProps } from "./ui/button";

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
                        providerName="Google"
                        href={`/api/v1/login/google`}
                        className="cursor-pointer text-[#db4639] hover:text-[#db4639]/90"
                    />
                </DropdownMenuItem>
                <DropdownMenuItem asChild>
                    <ProviderLoginBtn
                        className="cursor-pointer"
                        LogoIcon={FaGithub}
                        providerName="Github"
                        href={`/api/v1/login/github`}
                    />
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
