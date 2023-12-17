import Link from 'next/link'
import React from 'react'

const Navbar = () => {
  return (
    <nav className='flex'>
        <div className='group'>

        <Link href="/">GoDevs</Link>
        <div className='h-1 w-0 duration-150 group-hover:w-full transition-all bg-green-500'></div>
        </div>
    </nav>
  )
}

export default Navbar