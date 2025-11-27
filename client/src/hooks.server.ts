import { error, type Handle } from '@sveltejs/kit';
import { jwtVerify } from 'jose';
import { PERSONABOX_SECRET } from '$env/static/private';
import { JWTExpired } from 'jose/errors';

export const handle: Handle = async ({ event, resolve }) => {
	try {
		const accessToken = event.cookies.get('pb_access_token');

		if (!PERSONABOX_SECRET) {
			console.log('invalid secret key');
			event.locals.user = null;
            throw error(500, "missing secret key")
		}

		const secretKey = new TextEncoder().encode(PERSONABOX_SECRET);

		if (!accessToken) {
			console.log('Invalid access_token');
			event.locals.user = null;
			return resolve(event);
		}

		const { payload } = await jwtVerify(accessToken, secretKey, {clockTolerance:30});

		if (typeof payload.sub != 'string' || typeof payload.email != 'string') {
			console.log('invalid email or sub');
			event.locals.user = null;
			return resolve(event);
		}

		const user = {
			sub: payload.sub,
			email: payload.email
		};

		event.locals.user = user;

		return resolve(event);
	} catch (err) {

        console.log('JWT verification failed: ',{
            name: (err as any).name,
            code: (err as any).code,
            message: (err as any).message
        })

		event.locals.user = null;
		event.cookies.delete('pb_access_token', {
			path: '/'
		});
		return resolve(event);
	}
};
