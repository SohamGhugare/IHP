import React from 'react'
import Profile from "../../assets/profile.png"
import { FaHospitalAlt } from "react-icons/fa"

const About = () => {
    return (
        <div>

            <text className="text-7xl justify-self-start m-20 font-bold">
                About
            </text>
            <div className="flex flex-col items-center justify-center m-10">
                <div className="w-11/12 md:w-500 h-80 md:h-320 rounded-xl flex overflow-hidden">
                    <div className="w-281 h-281 rounded-full overflow-hidden">
                        <img src={Profile} alt="Profile" className="w-full h-full object-cover" />
                    </div>


                    <div className="flex flex-col justify-center flex-1 m-10">
                        <div className="flex flex-col justify-center m-5">

                            <text className="text-4xl font-bold text-gray-700">
                                Dr. Rahul Khatri
                            </text>
                            <text className="text-2xl text-gray-400">
                                Pediatric Cadiologist | Senior Consultant
                            </text>
                        </div>
                        <div className="flex flex-col justify-center m-5">
                            <text className="text-xl text-black">
                                Registration Number
                            </text>
                            <text className="text-2xl font-bold text-gray-400">
                                69789
                            </text>
                        </div>
                        <div className="flex flex-col justify-center m-5">
                            <div className="flex items-center">
                                <FaHospitalAlt className="w-20 h-20" />
                                <div className="flex flex-col justify-center m-5">
                                    <text className="text-xl text-black">
                                        Fortis
                                    </text>
                                    <text className="text-2xl text-gray-400">
                                        Location
                                    </text>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
            <div className="flex flex-col items-center justify-center m-20">
                <text className="text-2xl">
                Dr. Rahul Khatri, an ear, nose and throat specialist has practiced Otolaryngology for over 25 years. He has received training in the best Otolaryngological centers around the world. He has various publications to his credit and has been invited as a faculty at many courses and workshops. He himself has conducted training workshops on endoscopic sinus surgery, neonatal hearing screening, preoperative evaluation, surgery and postoperative rehabilitation for cochlear implants etc. Dr. Neeraj Kasliwal is running a successful Cochlear Implant program in Jaipur since 2007.
                </text>
            </div>
        </div>
    )
}

export default About