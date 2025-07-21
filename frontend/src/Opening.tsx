import React from 'react';
import './Opening.css';

function Opening() {
  return (
    <div>
      
      {/* Navbar */}
      <nav className="navbar">
        <div className="logo-box">Logo</div>
        <div className="navbar-menu">
          <span>Home</span>
          <span>Explore</span>
          <span>Tournaments</span>
          <span>Players</span>
        </div>
      </nav>

      {/* Main Section */}
      <div className="main-section">
        <div className="figma-filled-section">
          
          {/* Stage Image Grid — now placed behind */}
          <div className="stage-image-grid">
            {/* Row 1 — Only 2 images, Battlefield and Final Destination */}
            <div className="stage-row-carousel">
              <div className="carousel-track row1-track">
                <img src="/Battlefield.png" alt="Battlefield" />
                <img src="/FinalDestination.jpg" alt="Final Destination" />
                <img src="/Battlefield.png" alt="Battlefield duplicate" />
                <img src="/FinalDestination.jpg" alt="Final Destination duplicate" />
              </div>
            </div>


            {/* Row 2 */}
            <div className="row2-wrapper">
              <div className="stage-row-carousel row2-carousel">
                <div className="carousel-track row2-track">
                  <img src="/PokemonStadium2.png" alt="Pokemon Stadium 2" />
                  <img src="/Smashville.png" alt="Smashville" />
                  <img src="/PokemonStadium2.png" alt="Pokemon Stadium 2 duplicate" />
                  <img src="/Smashville.png" alt="Smashville duplicate" />
                </div>
              </div>
            </div>


            {/* Row 3 */}
            <div className="row3-wrapper">
              <div className="stage-row-carousel row3-carousel">
                <div className="carousel-track row3-track">
                  <img src="/TownAndCity.jpg" alt="Town and City" />
                  <img src="/HollowBastion.jpg" alt="Hollow Bastion" />
                  <img src="/TownAndCity.jpg" alt="Town and City duplicate" />
                  <img src="/HollowBastion.jpg" alt="Hollow Bastion duplicate" />
                </div>
              </div>
            </div>
          </div>


          {/* Foreground Content */}
          <div className="content-wrapper">
            <div className="circle-logo">Logo</div>

            <div className="text-button-section">

              <h1 className="hero-text">Smash Matchmaking: Simplified.</h1>

              <div className="button-row">

                <button className="button type1 alt-hover-purple">
                  <span className="btn-txt">Play Now</span>
                </button>

                <button className="button type1">
                  <span className="btn-txt">Explore</span>
                </button>

              </div>
              
            </div>

          </div>
        </div>
      </div>
    </div>
  );
}

export default Opening;
