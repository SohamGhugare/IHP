import React from 'react'
import { MdLocalPharmacy } from 'react-icons/md'

const TabPanel = () => {
    return (
        <div class="flex justify-evenly w-full bg-blue-200 h-[300px] p-5">
            <div class="w-1/5 h-full bg-white rounded-lg flex flex-col justify-center ">
                <div className="flex flex-col text-center">
                    <p className="ml-2 text-2xl text-black">Pharmacy</p>
                </div>
            </div>
            <div class="w-1/5 h-full bg-white rounded-lg flex flex-col justify-center ">
                <div className="flex flex-col text-center">
                    <p className="ml-2 text-2xl text-black">Diagnostic Labs</p>
                </div>
            </div>
            <div class="w-1/5 h-full bg-white rounded-lg flex flex-col justify-center ">
                <div className="flex flex-col text-center">
                    <p className="ml-2 text-2xl text-black">Emergency Beds</p>
                </div>
            </div>
        </div>
    )
}

export default TabPanel