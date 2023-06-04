import React from 'react'
import BannerImg from "../../assets/banner_user.png"

const Banner = () => {
    return (
        <div className="flex flex-col justify-between items-center justify-center">
            
            <div className="w-11/12 md:w-500 h-80 md:h-320 bg-gradient-to-br from-blue-900 via-blue-400 to-pink-500 rounded-xl flex overflow-hidden">
                <div className="flex flex-col justify-center flex-2 m-10">
                    <text className="text-7xl font-bold text-white">
                    Your Health pictures,
                    </text>
                    <text className="text-7xl font-bold text-white">
                    Simplified and Secured
                    </text>
                </div>
                <div className='w-2/12'>

                </div>
                <img className="w-372 h-341 flex-1 mr-5" src={BannerImg} alt="banner" />
            </div>
            
        </div>
    )
}

export default Banner