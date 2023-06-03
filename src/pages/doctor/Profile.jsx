import React from 'react'
import Header from '../../components/common/Header'
import Banner from '../../components/Doctor/Banner'
import Searchbar from '../../components/Doctor/Searchbar'
import About from '../../components/Doctor/About'
import { BsQrCodeScan } from "react-icons/bs";
import Fab from '@mui/material/Fab';



const Profile = () => {
  return (
    <div className="flex flex-col justify-between ">
      <Header />
      <Banner />
      <div class="fixed bottom-6 right-6">
        <button class="bg-blue-500 hover:bg-blue-600 text-white w-16 h-16 rounded-full flex items-center justify-center">
          <BsQrCodeScan className="w-6 h-6"/>
        </button>
      </div>
      <Searchbar />
      {/* <div>
      for popup table
      </div> */}

      <div className="mt-10">
        <About />
      </div>
    </div>
  )
}

export default Profile