import Link from "next/link";
import { CheckCircle, HelpCircle, ArrowRight, MessageCircle, Zap, Shield, Users } from "lucide-react";
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Pricing - RobusTest | Enterprise Mobile Device Lab Pricing",
  description: "Custom pricing for enterprise mobile device lab management. No per-minute billing, unlimited users, unlimited test minutes. Contact us for a quote.",
};

const includedFeatures = [
  "Manual testing with live device streaming",
  "Test automation (Appium, Espresso, XCUITest)",
  "Performance monitoring (CPU, memory, network)",
  "Smart TV testing (Roku, Apple TV)",
  "CI/CD integrations (Jenkins, GitHub Actions, etc.)",
  "Detailed reporting & analytics",
  "Unlimited users",
  "Unlimited test minutes",
  "On-site installation & training",
  "Ongoing support & updates",
  "Full API access",
  "Air-gapped deployment support",
];

const pricingBenefits = [
  {
    icon: Zap,
    title: "No Per-Minute Billing",
    description: "Unlike cloud services that charge $0.17+ per device minute, RobusTest offers predictable yearly pricing with unlimited usage.",
  },
  {
    icon: Users,
    title: "Unlimited Users",
    description: "Add as many team members as you need. No per-seat charges or user tiers.",
  },
  {
    icon: Shield,
    title: "Your Infrastructure",
    description: "Deploy on your premises with your security controls. No data ever leaves your network.",
  },
];

const faqs = [
  {
    question: "How does pricing work?",
    answer: "We offer custom yearly pricing based on your specific requirements — number of devices, platforms needed, and support level. Contact us for a detailed quote tailored to your needs.",
  },
  {
    question: "What hardware is included?",
    answer: "We provide specifications for server requirements and can assist with device rack setup and hardware procurement. The exact hardware depends on your deployment size and requirements.",
  },
  {
    question: "Do you offer a trial or pilot program?",
    answer: "Yes. We offer a pilot program for qualified enterprises. This includes a limited deployment so you can evaluate the platform with your actual workflows before committing.",
  },
  {
    question: "What support is included?",
    answer: "All deployments include software updates, bug fixes, and technical support. We also provide on-site installation, team training, and ongoing assistance.",
  },
  {
    question: "Can we use our existing devices?",
    answer: "Yes. RobusTest works with standard Android and iOS devices. We'll help you verify compatibility and set up your existing device inventory.",
  },
  {
    question: "What about compliance and security?",
    answer: "Since RobusTest is deployed on your premises, you maintain full control over security and compliance. The platform supports air-gapped environments and enterprise security requirements.",
  },
];

export default function PricingPage() {
  return (
    <div className="pt-16">
      {/* Hero Section */}
      <section className="bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 text-white py-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h1 className="text-4xl md:text-5xl font-bold mb-6">
            Custom Pricing for Your Needs
          </h1>
          <p className="text-xl text-gray-300 max-w-2xl mx-auto">
            Every enterprise is different. We&apos;ll work with you to create a solution that fits your requirements and budget.
          </p>
        </div>
      </section>

      {/* Main Pricing Card */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="bg-white rounded-2xl shadow-xl overflow-hidden">
            <div className="bg-blue-600 text-white p-8 text-center">
              <h2 className="text-3xl font-bold mb-2">Enterprise Device Lab</h2>
              <p className="text-blue-100">Complete on-premise solution with hardware, software & services</p>
            </div>
            <div className="p-8 md:p-12">
              <div className="text-center mb-10">
                <p className="text-gray-600 mb-4">Yearly pricing based on your requirements</p>
                <div className="flex items-center justify-center gap-2 text-4xl font-bold text-gray-900">
                  <MessageCircle className="text-blue-600" size={36} />
                  <span>Let&apos;s Talk</span>
                </div>
                <p className="text-gray-500 mt-2">Custom quote for your organization</p>
              </div>

              <div className="grid md:grid-cols-2 gap-4 mb-10">
                {includedFeatures.map((feature) => (
                  <div key={feature} className="flex items-start gap-3">
                    <CheckCircle className="text-green-500 flex-shrink-0 mt-0.5" size={20} />
                    <span className="text-gray-700">{feature}</span>
                  </div>
                ))}
              </div>

              <div className="flex flex-col sm:flex-row gap-4 justify-center">
                <Link
                  href="/contact"
                  className="inline-flex items-center justify-center gap-2 bg-blue-600 text-white px-8 py-4 rounded-lg font-semibold hover:bg-blue-700 transition-colors text-lg"
                >
                  Get a Custom Quote
                  <ArrowRight size={20} />
                </Link>
                <Link
                  href="/contact"
                  className="inline-flex items-center justify-center gap-2 border-2 border-gray-300 text-gray-700 px-8 py-4 rounded-lg font-semibold hover:border-gray-400 hover:bg-gray-50 transition-colors text-lg"
                >
                  Schedule a Call
                </Link>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Why Our Pricing */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Why Our Pricing Model Works
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Predictable costs without the surprises of per-minute billing.
            </p>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            {pricingBenefits.map((benefit) => (
              <div key={benefit.title} className="text-center">
                <div className="w-14 h-14 bg-blue-100 rounded-xl flex items-center justify-center mx-auto mb-4">
                  <benefit.icon className="text-blue-600" size={28} />
                </div>
                <h3 className="text-xl font-semibold text-gray-900 mb-2">{benefit.title}</h3>
                <p className="text-gray-600">{benefit.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Cost Comparison */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Cloud vs. On-Premise: The Cost Reality
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              See how per-minute billing adds up for enterprise teams.
            </p>
          </div>
          <div className="max-w-4xl mx-auto">
            <div className="bg-white rounded-2xl p-8 shadow-lg">
              <h3 className="text-xl font-semibold text-gray-900 mb-6 text-center">
                Example: 100 devices, 8 hours/day testing
              </h3>
              <div className="grid md:grid-cols-2 gap-8">
                <div className="bg-red-50 rounded-xl p-6 border border-red-100">
                  <p className="text-sm text-red-600 font-medium mb-2">Cloud Service (per-minute)</p>
                  <p className="text-3xl font-bold text-red-600 mb-2">~$300,000+</p>
                  <p className="text-sm text-gray-600 mb-4">per year at $0.17/min</p>
                  <ul className="space-y-2 text-sm text-gray-600">
                    <li>• 100 devices × 8 hrs × 60 min = 48,000 min/day</li>
                    <li>• 48,000 × 250 work days = 12M min/year</li>
                    <li>• Plus additional per-user fees</li>
                    <li>• Data leaves your network</li>
                  </ul>
                </div>
                <div className="bg-green-50 rounded-xl p-6 border border-green-100">
                  <p className="text-sm text-green-600 font-medium mb-2">RobusTest (yearly)</p>
                  <p className="text-3xl font-bold text-green-600 mb-2">Fixed Price</p>
                  <p className="text-sm text-gray-600 mb-4">predictable annual cost</p>
                  <ul className="space-y-2 text-sm text-gray-600">
                    <li>• Unlimited test minutes</li>
                    <li>• Unlimited users</li>
                    <li>• Data stays on your network</li>
                    <li>• Hardware + Software + Support included</li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* What's Included */}
      <section className="py-16 md:py-24 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              What&apos;s Included
            </h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              A complete solution, not just software.
            </p>
          </div>
          <div className="grid md:grid-cols-4 gap-6">
            {[
              { title: "Software License", description: "Full RobusTest platform with all features" },
              { title: "Hardware Setup", description: "Assistance with device racks, servers, and infrastructure" },
              { title: "Installation & Training", description: "On-site setup and hands-on team training" },
              { title: "Ongoing Support", description: "Updates, bug fixes, and technical assistance" },
            ].map((item) => (
              <div key={item.title} className="bg-gray-50 rounded-xl p-6">
                <h3 className="font-semibold text-gray-900 mb-2">{item.title}</h3>
                <p className="text-sm text-gray-600">{item.description}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* FAQs */}
      <section className="py-16 md:py-24 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-16">
            <h2 className="text-3xl md:text-4xl font-bold text-gray-900 mb-4">
              Frequently Asked Questions
            </h2>
          </div>
          <div className="max-w-3xl mx-auto">
            <div className="space-y-6">
              {faqs.map((faq) => (
                <div
                  key={faq.question}
                  className="bg-white rounded-xl p-6 shadow-sm"
                >
                  <div className="flex items-start gap-3">
                    <HelpCircle className="text-blue-600 flex-shrink-0 mt-1" size={20} />
                    <div>
                      <h3 className="font-semibold text-gray-900 mb-2">
                        {faq.question}
                      </h3>
                      <p className="text-gray-600">{faq.answer}</p>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-16 md:py-24 gradient-primary text-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <h2 className="text-3xl md:text-4xl font-bold mb-6">
            Ready to Get Started?
          </h2>
          <p className="text-xl text-blue-100 mb-8 max-w-2xl mx-auto">
            Tell us about your requirements and we&apos;ll provide a detailed proposal tailored to your needs.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link
              href="/contact"
              className="inline-flex items-center justify-center gap-2 bg-white text-blue-600 px-8 py-4 rounded-lg font-semibold hover:bg-gray-100 transition-colors text-lg"
            >
              Contact Us for Pricing
              <ArrowRight size={20} />
            </Link>
          </div>
        </div>
      </section>
    </div>
  );
}
