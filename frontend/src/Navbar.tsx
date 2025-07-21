import React from 'react'
import './Navbar.css';

function Navbar() {
  return (

        <nav className="navbar">
            <div className="logo-box">Logo</div>
                <div className="navbar-menu">
                <span>Home</span>
                <span>Explore</span>
                <span>Tournaments</span>
                <span>Players</span>
            </div>
        </nav>

  )
}

export default Navbar;
