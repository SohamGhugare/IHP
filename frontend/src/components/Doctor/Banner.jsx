import React from 'react'
import BannerImg from "../../assets/banner_doc.png"
import TextAnimation from './TextAnimation'

const Banner = () => {
    return (
        <div className="flex flex-col items-center justify-center">
            
            <div className="w-11/12 md:w-500 h-80 md:h-320 bg-gradient-to-r from-green-700 to-cyan-400 rounded-xl flex overflow-hidden">
                    <TextAnimation/>

                
                <img className="h-90 flex-1" src={BannerImg} alt="banner" />
            </div>
        </div>
    )
}

export default Banner