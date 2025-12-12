import Link from "next/link";
import {
  Shield,
  Server,
  Smartphone,
  Monitor,
  Zap,
  BarChart3,
  Settings,
  Users,
  CheckCircle,
  ArrowRight,
  Tv,
  Cpu,
  Lock,
  Cloud,
} from "lucide-react";

const platforms = [
  { name: "Android", icon: Smartphone },
  { name: "iOS", icon: Smartphone },
  { name: "Roku", icon: Tv },
  { name: "Apple TV", icon: Tv },
  { name: "Chromecast", icon: Tv },
];

const benefits = [
  {
    icon: Lock,
    title: "100% On-Premise",
    description: "Your test data never leaves your network. Complete control over your infrastructure.",
  },
  {
    icon: Shield,
    title: "Enterprise Security",
    description: "Meet SOC2, HIPAA, and PCI-DSS compliance requirements with air-gapped deployments.",
  },
  {
    icon: Zap,
    title: "No Usage Caps",
    description: "Unlimited users, unlimited test minutes. Predictable yearly pricing with no surprises.",
  },
];

const features = [
  {
    icon: Smartphone,
    title: "Manual Testing",
    description: "Live device streaming with real-time touch and keyboard input. Debug issues as they happen.",
  },
  {
    icon: Settings,
    title: "Test Automation",
    description: "Run Appium, Espresso, and XCUITest suites with parallel execution across your device fleet.",
  },
  {
    icon: BarChart3,
    title: "Performance Monitoring",
    description: "Built-in CPU, memory, and network metrics. No additional tools required.",
  },
  {
    icon: Tv,
    title: "Smart TV Testing",
    description: "Test on Roku, Apple TV, and Chromecast devices. The only platform with full TV support.",
  },
  {
    icon: Cpu,
    title: "CI/CD Integration",
    description: "Jenkins, GitHub Actions, GitLab CI, and more. Fits into your existing pipeline.",
  },
  {
    icon: Users,
    title: "Team Collaboration",
    description: "Role-based access, project management, and detailed reporting for distributed teams.",
  },
];

const howItWorks = [
  {
    step: "01",
    title: "Discovery",
    description: "We assess your testing requirements, device needs, and infrastructure setup.",
  },
  {
    step: "02",
    title: "Deployment",
    description: "Hardware and software installed on-site by our engineering team.",
  },
  {
    step: "03",
    title: "Training",
    description: "Your team gets hands-on training to maximize platform utilization.",
  },
  {
    step: "04",
    title: "Support",
    description: "Ongoing maintenance, updates, and technical support included.",
  },
];

const comparisonPoints = [
  { feature: "Data stays in your network", robustest: true, cloud: false },
  { feature: "No per-minute billing", robustest: true, cloud: false },
  { feature: "Unlimited users", robustest: true, cloud: false },
  { feature: "Air-gapped deployment", robustest: true, cloud: false },
  { feature: "Smart TV testing", robustest: true, cloud: false },
  { feature: "Hardware included", robustest: true, cloud: false },
];

export default function Home() {
  return (
    <div className="pt-16">
      {/* Hero Section */}
      <section className="relative bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-white overflow-hidden">
        <div className="absolute inset-0 bg-[url('/grid.svg')] opacity-10"></div>
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20 md:py-32 relative">
          <div className="max-w-3xl">
            <h1 className="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight mb-6">
              Stop Managing Devices.<br />
              <span className="text-blue-400">Start Testing.</span>
            </h1>
            <p className="text-xl md:text-2xl text-gray-300 mb-8 leading-relaxed">
              We set up and manage your in-house device lab so your team can focus on quality. Hardware, software, and expertise — delivered to your premises.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 mb-12">
              <Link
                href="/contact"
                className="bg-blue-600 text-white px-8 py-4 rounded-lg font-semibold hover:bg-blue-700 transition-colors text-center text-lg"
              >
                Request a Demo
              </Link>
              <Link
                href="/features"
                className="border-2 border-gray-600 text-white px-8 py-4 rounded-lg font-semibold hover:border-gray-400 hover:bg-gray-800 transition-colors text-center text-lg"
              >
                Explore Features
              </Link>
            </div>
            <div className="flex flex-wrap items-center gap-6 text-sm text-gray-400">
              <div className="flex items-center gap-2">
                <CheckCircle size={18} className="text-green-500" />
                <span>100% On-Premise</span>
              </div>
              <div className="flex items-center gap-2">
                <CheckCircle size={18} className="text-green-500" />
                <span>No Per-Minute Billing</span>
              </div>
              <div className="flex items-center gap-2">
                <CheckCircle size={18} className="text-green-500" />
                <span>Enterprise Support</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Platforms Section */}
      <section className="py-12 bg-gray-50 border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <p className="text-center text-gray-500 mb-8 font-medium">SUPPORTED PLATFORMS</p>
          <div className="flex flex-wrap justify-center items-center gap-8 md:gap-12">
            {platforms.map((platform) => (
              <div key={platform.name} className="flex items-center gap-2 text-gray-600">
                <platform.icon size={24} />
                <span className="font-medium">{platform.name}</span>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Why On-Premise Section */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Why Enterprise Teams Choose On-Premise
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Cloud testing services work for some. But when security, compliance, and cost predictability matter, on-premise is the answer.
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            {benefits.map((benefit) => (
              <div
                key={benefit.title}
                className="bg-gray-50 rounded-xl p-8 card-hover"
              >
                <div className="w-12 h-12 bg-blue-100 rounded-lg flex items-center justify-center mb-6">
                  <benefit.icon className="text-blue-600" size={24} />
                </div>
                <h3 className="text-xl font-semibold text-gray-900 mb-3">
                  {benefit.title}
                </h3>
                <p className="text-gray-600 leading-relaxed">
                  {benefit.description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* What We Deliver Section */}
      <section className="py-16 md:py-24 bg-gray-900 text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold mb-4">
              Everything You Need to Run a Device Lab
            </h2>
            <p className="text-xl text-gray-400 max-w-2xl mx-auto">
              We deliver a complete solution — not just software.
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            <div className="bg-gray-800 rounded-xl p-8">
              <Server className="text-blue-400 mb-6" size={40} />
              <h3 className="text-xl font-semibold mb-4">Hardware</h3>
              <ul className="space-y-3 text-gray-400">
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Device racks and mounting</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>USB hubs and connectivity</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Power management systems</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Network infrastructure</span>
                </li>
              </ul>
            </div>
            <div className="bg-gray-800 rounded-xl p-8">
              <Monitor className="text-blue-400 mb-6" size={40} />
              <h3 className="text-xl font-semibold mb-4">Software</h3>
              <ul className="space-y-3 text-gray-400">
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>RobusTest platform license</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Full API access</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>CI/CD plugins</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Regular updates</span>
                </li>
              </ul>
            </div>
            <div className="bg-gray-800 rounded-xl p-8">
              <Users className="text-blue-400 mb-6" size={40} />
              <h3 className="text-xl font-semibold mb-4">Services</h3>
              <ul className="space-y-3 text-gray-400">
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>On-site installation</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Team training</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Ongoing support</span>
                </li>
                <li className="flex items-start gap-2">
                  <CheckCircle size={18} className="text-green-500 mt-1 flex-shrink-0" />
                  <span>Device procurement assistance</span>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Platform Capabilities
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Everything your team needs to test mobile and TV applications effectively.
            </p>
          </div>
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {features.map((feature) => (
              <div
                key={feature.title}
                className="border border-gray-200 rounded-xl p-6 card-hover"
              >
                <div className="w-10 h-10 bg-blue-100 rounded-lg flex items-center justify-center mb-4">
                  <feature.icon className="text-blue-600" size={20} />
                </div>
                <h3 className="text-lg font-semibold text-gray-900 mb-2">
                  {feature.title}
                </h3>
                <p className="text-gray-600">
                  {feature.description}
                </p>
              </div>
            ))}
          </div>
          <div className="text-center mt-12">
            <Link
              href="/features"
              className="inline-flex items-center gap-2 text-blue-600 font-semibold hover:text-blue-700 transition-colors"
            >
              See all features
              <ArrowRight size={18} />
            </Link>
          </div>
        </div>
      </section>

      {/* Comparison Section */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              RobusTest vs. Cloud Testing Services
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              See why enterprises are moving from BrowserStack, Sauce Labs, and AWS Device Farm to on-premise solutions.
            </p>
          </div>
          <div className="max-w-3xl mx-auto">
            <div className="bg-white rounded-xl shadow-lg overflow-hidden">
              <div className="grid grid-cols-3 bg-gray-100 p-4 font-semibold text-gray-700">
                <div>Feature</div>
                <div className="text-center">RobusTest</div>
                <div className="text-center">Cloud Services</div>
              </div>
              {comparisonPoints.map((point, index) => (
                <div
                  key={point.feature}
                  className={`grid grid-cols-3 p-4 ${
                    index % 2 === 0 ? "bg-white" : "bg-gray-50"
                  }`}
                >
                  <div className="text-gray-700">{point.feature}</div>
                  <div className="text-center">
                    {point.robustest ? (
                      <CheckCircle className="inline text-green-500" size={20} />
                    ) : (
                      <span className="text-gray-400">-</span>
                    )}
                  </div>
                  <div className="text-center">
                    {point.cloud ? (
                      <CheckCircle className="inline text-green-500" size={20} />
                    ) : (
                      <span className="text-red-400">&#10005;</span>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* How It Works Section */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              How It Works
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              From initial consultation to ongoing support, we handle everything.
            </p>
          </div>
          <div className="grid md:grid-cols-4 gap-8">
            {howItWorks.map((step, index) => (
              <div key={step.step} className="relative">
                <div className="text-5xl font-bold text-blue-100 mb-4">
                  {step.step}
                </div>
                <h3 className="text-xl font-semibold text-gray-900 mb-2">
                  {step.title}
                </h3>
                <p className="text-gray-600">{step.description}</p>
                {index < howItWorks.length - 1 && (
                  <div className="hidden md:block absolute top-8 right-0 w-1/2 h-0.5 bg-blue-100"></div>
                )}
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Industries Section */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Trusted by Enterprise Teams
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Deployed in device labs across industries where security and quality matter.
            </p>
          </div>
          <div className="grid md:grid-cols-4 gap-6 text-center">
            {["E-Commerce", "Streaming & Media", "Banking & Fintech", "Healthcare"].map(
              (industry) => (
                <div
                  key={industry}
                  className="bg-white rounded-xl p-6 shadow-sm"
                >
                  <p className="font-semibold text-gray-700">{industry}</p>
                </div>
              )
            )}
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 md:py-24 gradient-primary text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl md:text-4xl font-bold mb-6">
            Ready to Transform Your Device Lab?
          </h2>
          <p className="text-xl text-blue-100 mb-8 max-w-2xl mx-auto">
            Schedule a consultation with our team to discuss your requirements and see RobusTest in action.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link
              href="/contact"
              className="bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition-colors text-lg"
            >
              Schedule a Consultation
            </Link>
            <Link
              href="/pricing"
              className="border-2 border-white text-white px-8 py-4 rounded-lg font-semibold hover:bg-white/10 transition-colors text-lg"
            >
              View Pricing
            </Link>
          </div>
        </div>
      </section>
    </div>
  );
}
