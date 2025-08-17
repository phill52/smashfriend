import "./profile.css"

function ProfileComponent({imageUrl, userName, userId, rank, elo_rating, win_rate, loss_rate}) {
    return (
        <div className="main-container">
            <div className="profile-component">
                Your Profile
                <div className="main-content">
                    <img src={imageUrl} alt="user profile"/>
                    <div className="profile-data">
                        <p>{userName}</p>
                        <div className="api-information">
                            <h6>Start gg ID: {userId}</h6>
                        </div>
                        <div className="rank-information">
                            <p>Current Rank</p>
                            <div className="triangle-down"></div>
                            <p>{rank}</p>
                            <div className="api-information">
                                <h6>{elo_rating} ELO</h6>
                            </div>
                            <p>This Season:</p>
                            <div id="win-rate">
                                <p><span style={{color: "green"}}>{win_rate} W</span>-<span style={{color: "red"}}> {loss_rate} L </span>(74%)</p>
                            </div>
                        </div>
                        <div style={{margin: "16px"}}>
                            <button id="profile-button">Full Profile</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default ProfileComponent;