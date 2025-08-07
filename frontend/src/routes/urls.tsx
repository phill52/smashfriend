import { registerRoutes } from "./routeBuilder";

export const routes = registerRoutes({
    index: "/",
    login: "/login",
    lobby: {
        lobby: "/lobby",
    },
    match: {
        match: "/match/:id",
    },
});
