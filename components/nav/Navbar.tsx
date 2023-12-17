import Link from "next/link";
import React from "react";
import { Button } from "../ui/button";
import { SiGithub } from "react-icons/si";

const Navbar = () => {
  return (
    <nav className="flex items-center justify-between">
      <div className="group">
        <Link href="/" className="font-bold text-[2rem]">
          <span className="font-semibold">GoDevs</span>
        </Link>
        <div className="h-1 w-0 duration-150 group-hover:w-full transition-all bg-blue-500" />
      </div>

      <Button variant="outline" className="flex items-center gap-2">
        <SiGithub /> Login
      </Button>
    </nav>
  );
};

export default Navbar;
