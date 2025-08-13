import "./profile.css"

function ProfileComponent() {
    return (
        <div className="main-container">
            <div className="profile-component">
                Your Profile
                <div className="main-content">
                    <div className="img-placeholder"></div>
                    <div className="profile-data">
                        <p>dan</p>
                        <div className="api-information">
                            <h6>Start gg ID: 123456789</h6>
                        </div>
                        <div className="rank-information">
                            <p>Current Rank</p>
                            <div className="triangle-down"></div>
                            <p>Universal</p>
                            <div className="api-information">
                                <h6>2500 ELO</h6>
                            </div>
                            <p>This Season:</p>
                            <div id="win-rate">
                                <p><span style={{color: "green"}}>152 W</span>-<span style={{color: "red"}}> 53 L</span>(74%)</p>
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