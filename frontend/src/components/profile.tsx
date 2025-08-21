import "./profile.css";

function ProfileComponent({ user }) {
  const win_rate = Math.floor(
    (user.games_won / (user.games_won + user.games_lost)) * 100,
  );
  return (
    <div className="flex min-h-[calc(100vh-4.75rem)] w-full justify-end text-white">
      <div className="mt-1.5 mr-2 rounded-lg border border-[#38334E] bg-[#191E29] p-2.5 text-center font-bold text-white">
        Your Profile
        <div className="p-4">
          <img
            className="mx-auto h-24 w-24 rounded-full bg-white"
            src={user.image_url}
            alt="user profile"
          />
          <div className="pt-1.5 text-sm">
            <p>{user.username}</p>
            <div className="mb-0.5 text-xs text-[#CACACA]">
              <h6>Start gg ID: {user.user_id}</h6>
            </div>
            <div className="pt-1">
              <p className="m-0.5">Current Rank</p>
              <div className="m-1.5 mx-auto h-0 w-0 border-t-56 border-r-32 border-l-32 border-t-white border-r-transparent border-l-transparent"></div>
              <p>{user.rank}</p>
              <div className="mb-1 text-xs text-[#CACACA]">
                <h6>{user.elo_rating} ELO</h6>
              </div>
              <p>This Season:</p>
              <div className="mt-1 text-sm whitespace-nowrap">
                <p>
                  <span className="text-green-600">{user.games_won} W</span>-
                  <span className="text-red-600"> {user.games_lost} L </span>(
                  {win_rate}%)
                </p>
              </div>
            </div>
            <div>
              <button className="mt-2 cursor-pointer rounded-lg bg-[#96619A] px-3 py-2 text-sm font-bold">
                Full Profile &gt;
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default ProfileComponent;
