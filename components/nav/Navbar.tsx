"use client";
import Link from "next/link";
import React from "react";
import Login from "./LoginForm";
import { useUser } from "@/lib/store/user";

const Navbar = () => {
  const user = useUser((state) => state.users);
  return (
    <nav className="flex items-center justify-between">
      <div className="group">
        <Link href="/" className="font-bold text-[2rem]">
          <span className="font-semibold">GoDevs</span>
        </Link>
        <div className="h-1 w-0 duration-150 group-hover:w-full transition-all bg-blue-500" />
      </div>

      {user?.id ? <h1>Profile</h1> : <Login />}
    </nav>
  );
};

export default Navbar;
