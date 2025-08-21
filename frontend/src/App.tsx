import {
  SignedIn,
  SignInButton,
  SignedOut,
  SignOutButton,
  UserButton,
} from "@clerk/clerk-react";

import "./index.css";
import "./App.css";
import Profile from "./components/profile";
import react_path from "./assets/react.svg";

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
  return (
    <>
      <>
        <Profile user={user} />
      </>
      <SignedOut>
        <SignInButton />
      </SignedOut>
      <SignedIn>
        <UserButton />
        <SignOutButton />
      </SignedIn>
    </>
  );
}

export default App;
