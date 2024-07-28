"use client";

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

export default function Dashboard() {
  const router = useRouter();
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    // 仮の認証チェック
    const user = localStorage.getItem('user');
    if (user) {
      setIsAuthenticated(true);
    } else {
      // ログインしていなくても一時的に閲覧可能にするため、リダイレクトをコメントアウト
      // router.push('/signin');
    }
  }, [router]);

  const handleSignOut = () => {
    localStorage.removeItem('user');
    setIsAuthenticated(false);
    router.push('/signin');
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">
      <div className="bg-gray-800 p-8 rounded shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-4">Dashboard</h2>
        <p>Welcome to your dashboard!</p>
        <button
          onClick={handleSignOut}
          className="mt-4 px-4 py-2 bg-red-500 text-white rounded"
        >
          Sign Out
        </button>
      </div>
    </div>
  );
}
