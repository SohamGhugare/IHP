import React from 'react'
import Header from "../../components/common/Header"
import Banner from '../../components/Patient/Banner'
import TabPanel from '../../components/Patient/TabPanel'
import Searchbar from '../../components/Patient/Searchbar'
import TextHead from '../../components/Patient/TextHead'
import Footer from '../../components/Patient/Footer'
import { Link } from 'react-router-dom'


const ProfileUser = () => {
  return (
    <div className="flex flex-col justify-between ">
      <Header />
      <Banner />
      <div class="fixed bottom-20 right-6">
        <Link to="https://sourceai.in/map">
          <button class="bg-red-500 text-white w-16 h-16 rounded-full flex items-center justify-center">
            SOS
          </button>
        </Link>
      </div>
      <div className="mt-10">
        <Searchbar />
      </div>
      <div className="my-10">
        <TextHead />
      </div>

      {/* <div>
      for popup table
      </div> */}

      <div className="m-10">
        <TabPanel/>
      </div>
      <Footer/>
    </div>
  )
}

export default ProfileUser