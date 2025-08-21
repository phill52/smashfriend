import {
  SignedIn,
  SignInButton,
  SignedOut,
  SignOutButton,
  UserButton,
  useAuth,
} from "@clerk/clerk-react";
import { useState } from "react";
import "./index.css";
import "./App.css";
import Profile from "./components/profile";
import react_path from "./assets/react.svg";
import viteLogo from "/vite.svg";
import reactLogo from "./assets/react.svg";

const user = {
  user_id: "123456789",
  image_url: react_path,
  username: "dan",
  rank: "Universal",
  elo_rating: "2500",
  games_won: 152,
  games_lost: 53,
};

function App() {
  const [count, setCount] = useState(0);
  const { getToken } = useAuth();

  return (
    <>
      <Profile user={user} />
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p className="font-bold text-red-500 underline">
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>

      {/* Clerk authentication (both versions had this) */}
      <SignedOut>
        <SignInButton />
      </SignedOut>
      <SignedIn>
        <UserButton />
        <SignOutButton />
        <button
          onClick={async () => {
            const token = await getToken();
            await fetch("http://localhost:8080/api/users", {
              method: "GET",
              headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
              },
            })
              .then((res) => res.json())
              .then((data) => {
                console.log(data);
              })
              .catch((err) => {
                console.error(err);
              });
          }}
          className="rounded-md bg-blue-500 p-2 text-white"
        >
          Ping backend
        </button>
      </SignedIn>
    </>
  );
}

export default App;
