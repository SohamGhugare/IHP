import React from 'react'
import { FaHospitalAlt } from "react-icons/fa"

const Experience = () => {
    return (
        <div>

            <text className="text-7xl justify-self-start m-20 font-bold">
                Experience
            </text>
            <div className="flex flex-col justify-center m-10">
                <div className="flex flex-col justify-center ml-10">

                    <div className="flex items-center">
                        <FaHospitalAlt className="w-10 h-10" />
                        <div className="flex flex-col justify-center m-5">
                            <text className="text-xl text-black">
                                Senior Consultant
                            </text>
                            <text className="text-2xl text-gray-400">
                                Medical Affairs
                            </text>
                        </div>
                    </div>
                    <div className="flex items-center">
                        <FaHospitalAlt className="w-10 h-10" />
                        <div className="flex flex-col justify-center m-5">
                            <text className="text-xl text-black">
                                Chief Otolaryngologist
                            </text>
                            <text className="text-2xl text-gray-400">
                                involved in specialist care of all major otolaryngological problems
                            </text>
                        </div>
                    </div>
                </div>
            </div>




        </div>
    )
}

export default Experience