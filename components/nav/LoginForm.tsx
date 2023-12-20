"use client";

import React from "react";
import { Button } from "../ui/button";
import { SiGithub } from "react-icons/si";
import { createBrowserClient } from "@supabase/ssr";
import { usePathname } from "next/navigation";

const Login = () => {
  const pathname = usePathname();

  const supabase = createBrowserClient(
    process.env.NEXT_PUBLIC_SUPABASE_URL!,
    process.env.NEXT_PUBLIC_SUPABASE_ANON_KEY!
  );

  const handleLogin = () => {
    supabase.auth.signInWithOAuth({
      provider: "github",
      options: {
        redirectTo: location.origin + "/auth/callback?next=" + pathname,
      },
    });
  };

  return (
    <Button
      variant="outline"
      className="flex items-center gap-2"
      onClick={handleLogin}
    >
      <SiGithub /> Login
    </Button>
  );
};

export default Login;
