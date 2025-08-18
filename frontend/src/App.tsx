// import {
//     SignedIn,
//     SignInButton,
//     SignedOut,
//     SignOutButton,
//     UserButton,
// } from "@clerk/clerk-react";

//import "./index.css"
import "./App.css";
import Profile from "./components/profile"
import reactPath from './assets/react.svg'

function App() {
    return (
        <div>
            <Profile 
                imageUrl={reactPath}
                userName="dan"
                userId="123456789"
                rank="Universal"
                elo_rating="2500"
                games_won={152}
                games_lost={53}
            />
        </div>

            // {/* <SignedOut>
            //     <SignInButton />
            // </SignedOut>
            // <SignedIn>
            //     <UserButton />
            //     <SignOutButton />
            // </SignedIn> */}
    );
}

export default App;
