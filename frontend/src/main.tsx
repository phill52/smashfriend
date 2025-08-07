import { RouterProvider } from "react-router";
import React from "react";
import ReactDOM from "react-dom/client";

import { ClerkProvider } from "@clerk/clerk-react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import router from "./router";

const queryClient = new QueryClient();

const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY;
if (!PUBLISHABLE_KEY) {
    throw new Error("Missing Publishable Key");
}

const root = document.getElementById("root")!;
ReactDOM.createRoot(root).render(
    <React.StrictMode>
        <QueryClientProvider client={queryClient}>
            <ClerkProvider publishableKey={PUBLISHABLE_KEY}>
                <RouterProvider router={router} />
            </ClerkProvider>
        </QueryClientProvider>
    </React.StrictMode>
);
