import "./profile.css"

function ProfileComponent({imageUrl, userName, userId, rank, elo_rating, games_won, games_lost}) {
    const win_rate = Math.floor((games_won / (games_won + games_lost)) * 100)
    return (
        <div className="flex justify-end items-center min-h-[calc(100vh-4.75rem)] text-white mr-[0.625rem]">
            <div className="w-[10.9375rem] h-[28.125rem] border border-[#38334E] text-center bg-[#191E29] rounded-[0.625rem] font-bold p-[0.625rem] text-white mr-[0.625rem]">
                Your Profile
                <div className="p-4">
                    <img className="rounded-full bg-white h-[6.25rem] w-[6.25rem] mx-auto" src={imageUrl} alt="user profile"/>
                    <div className="pt-[0.3125rem] text-[0.875rem]">
                        <p>{userName}</p>
                        <div className="text-[0.625rem] w-full h-full text-[#CACACA] mb-[0.3125rem]">
                            <h6>Start gg ID: {userId}</h6>
                        </div>
                        <div className="pt-[0.625rem]">
                            <p className="mb-[0.625rem]">Current Rank</p>
                            <div className="w-0 h-0 border-l-[2rem] border-r-[1.90625rem] border-t-[3.8125rem] border-l-transparent border-r-transparent border-t-white mx-auto"></div>
                            <p>{rank}</p>
                            <div className="text-[0.625rem] w-full h-full text-[#CACACA] mb-[0.3125rem]">
                                <h6>{elo_rating} ELO</h6>
                            </div>
                            <p>This Season:</p>
                            <div className="mt-[0.3125rem] whitespace-nowrap text-[0.9375rem]">
                                <p><span style={{color: "green"}}>{games_won} W</span>-<span style={{color: "red"}}> {games_lost} L </span>({win_rate}%)</p>
                            </div>
                        </div>
                        <div style={{margin: "1rem"}}>
                            <button className="bg-[#96619A] rounded-[0.625rem] text-[0.625rem] gap-[0.5rem] font-bold py-[0.5rem] px-[0.75rem] cursor-pointer">Full Profile &gt;</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default ProfileComponent;