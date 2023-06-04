import React from 'react'
import Header from '../../components/common/Header'
import Banner from '../../components/Doctor/Banner'
import Searchbar from '../../components/Doctor/Searchbar'
import About from '../../components/Doctor/About'
import { BsQrCodeScan } from "react-icons/bs";
import { Link } from 'react-router-dom'
import Experience from '../../components/Doctor/Experience'
import Awards from '../../components/Doctor/Awards'
import Video from '../../components/Doctor/Video'
import Footer from '../../components/Doctor/Footer'

const Profile = () => {
  return (
    <div className="flex flex-col justify-between ">
      <Header />
      <Banner />
      <div class="fixed bottom-20 right-6">
        <Link to="https://sourceai.in/qrscan">
          <button class="bg-blue-500 hover:bg-blue-600 text-white w-16 h-16 rounded-full flex items-center justify-center">
            <BsQrCodeScan className="w-6 h-6" />
          </button>
        </Link>
      </div>
      <Searchbar />
      {/* <div>
      for popup table
      </div> */}

      <div className="mt-10">
        <About />
      </div>
      
        <Experience/>
        <Awards/>
        <Video/>
        <Footer/>
      
    </div>
  )
}

export default Profile