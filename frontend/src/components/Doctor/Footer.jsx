import React from 'react'
import {BiLogOut} from 'react-icons/bi'
import Logo from "../../assets/full_logo.png"

const Footer = () => {
  return (
    <div>
        <footer className="flex items-center justify-between p-4 bg-green-700 text-white">
        <div className="flex items-center">
        {/* Logo */}
        <img src={Logo} alt="Logo" className="h-15" />
      </div>

      <div className="flex items-center">
        <button className='bg-teal-500 flex items-center justify-center rounded-xl py-5 px-5'>
        <BiLogOut className='m-1'/>
            sdvmsv
        </button>
      </div>
        </footer>
        
    </div>
  )
}

export default Footer