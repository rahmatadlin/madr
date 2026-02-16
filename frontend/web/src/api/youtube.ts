import { apiClient, type ApiResponse } from './client'

export interface YouTubeVideo {
  video_id: string
  title: string
  description: string
  published_at: string
  thumbnail_url: string
  channel_title: string
}

export interface YouTubeVideosResponse {
  data: YouTubeVideo[]
  total: number
}

export const youtubeApi = {
  getKajianVideos: async (): Promise<YouTubeVideosResponse> => {
    const { data } = await apiClient.get<YouTubeVideosResponse>('/youtube/kajian')
    return data
  },
}
