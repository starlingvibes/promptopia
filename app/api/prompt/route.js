import { connectToDB } from '@utils/database';
import Prompt from '@models/prompt';

export const GET = async () => {
  try {
    await connectToDB();
    const prompts = await Prompt.find({}).populate('creator');
    return new Response(JSON.stringify(prompts), { status: 200 });
  } catch (error) {
    return new Response(JSON.stringify('Failed to get prompts'), {
      status: 500,
    });
  }
};
