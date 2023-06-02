
const Roles = () => {
    return (
        <div className="flex flex-col items-center space-y-4">
            <text className="text-6xl font-bold text-white">
                Select Roles
            </text>
            <div className="flex space-x-5">
                <button variant="contained" className="py-2 px-4 text-lg bg-white text-black shadow-md rounded-xl">
                    Patient
                </button>
                <button variant="contained" className="py-2 px-4 text-lg bg-white text-black shadow-md rounded-xl">
                    Doctor
                </button>
            </div>
            
        </div>
    )
}

export default Roles