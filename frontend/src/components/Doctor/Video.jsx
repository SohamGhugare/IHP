import React from 'react'
import ReactPlayer from "react-player/youtube";

const Video = () => {
    return (
        <div>
            <text className="text-7xl justify-self-start m-20 font-bold">
                Video
            </text>
            <div className="flex flex-col justify-center mx-20 my-5">

                <div className="w-11/12 md:w-200 h-80 md:h-1000">
                    <ReactPlayer
                        url=""
                        controls
                        width="100%"
                        height="100%"
                        style={{ backgroundColor: "#000000" }}
                        playing={true}
                    />
                </div>
            </div>
        </div>
    )
}

export default Video