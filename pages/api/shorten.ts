import { NextApiRequest, NextApiResponse } from 'next';
import axios from 'axios';

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  if (req.method === 'POST') {
    const { url } = req.body;

    try {
      const backendResponse = await axios.post('http://localhost:8080/api/urls', {
        original_url: url,
      });

      if (backendResponse.status === 200) {
        const data = backendResponse.data;
        res.status(200).json({ shortenedUrl: data.short_code });
      } else {
        res.status(500).json({ error: 'Error shortening URL' });
      }
    } catch (error) {
      console.error('Error connecting to backend:', error);
      res.status(500).json({ error: 'Error shortening URL' });
    }
  } else {
    res.status(405).end();
  }
}