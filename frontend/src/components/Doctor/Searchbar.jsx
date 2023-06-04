import React from 'react'
import TextField from '@mui/material/TextField';
import { FiSearch, FiFilter } from 'react-icons/fi'

const Searchbar = () => {
    return (
        <div className="flex flex-col items-center justify-center">
            <div className="flex justify-evenly w-11/12 md:w-500 h-20 bg-gradient-to-r from-green-700 to-cyan-400 rounded-full m-5">
                <div className="flex items-center justify-center m-10 bg-gray-200 bg-opacity-50 rounded-full">
                    <FiSearch className="text-white text-2xl" />
                    <input
                        type="text"
                        className="w-full px-4 py-2 focus:outline-none bg-transparent border-b-2 border-white-500 text-white h-12 text-lg"
                    />
                </div>
                <div className="flex items-center justify-center m-10 bg-gray-200 bg-opacity-50 rounded-full">
                    <FiFilter className="text-white text-2xl" />
                    <input
                        type="text"
                        className="w-full px-4 py-2 focus:outline-none bg-transparent border-b-2 border-white-500 text-white h-12 text-lg"
                    />
                </div>
                <div className="flex items-center justify-center m-10 ">
                    <button className="flex items-center justify-center rounded-full bg-white py-5 px-10 ">
                        View Profile
                    </button>
                </div>
            </div>
        </div>
    )
}

export default Searchbar