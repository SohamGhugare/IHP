import React from 'react'
import { MdLocalPharmacy } from 'react-icons/md'

const TabPanel = () => {
    return (
        <div className="flex flex-col items-center justify-center bg-blue-100 h-200">
            <div className="carousel rounded-box flex items-center justify-center">
                <div className="carousel-item w-1000">
                    <div className="card w-90 bg-base-100 shadow-xl">
                        <figure className="px-10 pt-10">
                            <div className="card-actions">

                                <div className="w-20 h-20 rounded-full bg-blue-500 flex items-center justify-center">
                                    <MdLocalPharmacy className='w-10 h-10 text-white'/>
                                </div>
                            </div>
                        </figure>
                        <div className="card-body items-center text-center">
                            <h2 className="card-title">Shoes!</h2>
                            <p>If a dog chews shoes whose shoes does he choose?</p>

                        </div>
                    </div>
                </div>
                <div className="carousel-item w-1000">
                    <div className="card w-90 bg-base-100 shadow-xl">
                        <figure className="px-10 pt-10">
                            <img src="/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg" alt="Shoes" className="rounded-xl" />
                        </figure>
                        <div className="card-body items-center text-center">
                            <h2 className="card-title">Shoes!</h2>
                            <p>If a dog chews shoes whose shoes does he choose?</p>
                            <div className="card-actions">
                                <button className="btn btn-primary">Buy Now</button>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="carousel-item w-1000">
                    <div className="card w-90 bg-base-100 shadow-xl">
                        <figure className="px-10 pt-10">
                            <img src="/images/stock/photo-1606107557195-0e29a4b5b4aa.jpg" alt="Shoes" className="rounded-xl" />
                        </figure>
                        <div className="card-body items-center text-center">
                            <h2 className="card-title">Shoes!</h2>
                            <p>If a dog chews shoes whose shoes does he choose?</p>
                            <div className="card-actions">
                                <button className="btn btn-primary">Buy Now</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default TabPanel