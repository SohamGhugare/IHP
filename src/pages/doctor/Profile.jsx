import React from 'react'
import Header from '../../components/common/Header'
import Banner from '../../components/Doctor/Banner'
import Searchbar from '../../components/Doctor/Searchbar'

const Profile = () => {
  return (
    <div className="flex flex-col justify-between">
      <Header />
      <Banner />
      <div className="mt-10">
        <Searchbar />
      </div>
      {/* <div>
      for popup table
      </div> */}

      
    </div>
  )
}

export default Profile