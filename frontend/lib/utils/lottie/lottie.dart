import 'package:lottie/lottie.dart';

class LottieAnim {
  static LottieBuilder ForgotPassword(double width, double height) {
    return Lottie.asset('assets/lottie/forgot_password',
        width: width, height: height);
  }

  static LottieBuilder LoginOrRegister(double width, double height) {
    return Lottie.asset('assets/lottie/login_anim.json',
        width: width, height: height);
  }
}