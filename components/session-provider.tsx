"use client"

import { useUser } from "@/lib/store/user";
import { createBrowserClient } from "@supabase/ssr";
import React, { useEffect } from "react";

export default function SessionProvider() {
  const setUser = useUser((state) => state.setUser);

  const supabase = createBrowserClient(
    "https://omwpfturumksmfbrhteq.supabase.co",
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im9td3BmdHVydW1rc21mYnJodGVxIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDI5NzAyNzIsImV4cCI6MjAxODU0NjI3Mn0.JPknsQLa2__IrnuUK4cziuGyT-zbnqRfZlrGyFgn8pA"
  );

  const readUserSession = async () => {
    const { data } = await supabase.auth.getSession();
    setUser(data.session?.user);
  };

  useEffect(() => {
    readUserSession();
    //eslint-disable-next-line
  }, []);

  return <div></div>;
}
