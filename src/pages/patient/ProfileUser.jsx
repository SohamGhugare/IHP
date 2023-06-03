import React from 'react'
import Header from "../../components/common/Header"
import Banner from '../../components/Patient/Banner'
import TabPanel from '../../components/Patient/TabPanel'
import Searchbar from '../../components/Patient/Searchbar'
import TextHead from '../../components/Patient/TextHead'
import Footer from '../../components/Patient/Footer'

const ProfileUser = () => {
  return (
    <div className="flex flex-col justify-between ">
      <Header />
      <Banner />
      <div className="mt-10">
        <Searchbar />
      </div>
      <div className="mt-10">
        <TextHead />
      </div>

      {/* <div>
      for popup table
      </div> */}
{/* 
      <div className="mt-10">
        <TabPanel/>
      </div> */}
      <Footer/>
    </div>
  )
}

export default ProfileUser